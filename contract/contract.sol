// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

/**
 * @title Scraper
 * @dev A contract to manage decentralized scraping tasks. This version is highly
 * optimized for gas efficiency by minimizing on-chain storage.
 */
contract Scraper {
    address public owner;
    bool public paused;

    enum Status {
        Pending,
        Completed,
        Failed
    }

    // Using bytes32 for the mapping key is much more gas-efficient than string.
    mapping(bytes32 => Task) public tasks;

    struct Task {
        address payable submitter; // 20 bytes
        address payable assignee;
        Status status; // 1 byte
        uint256 amount; // 32 bytes
        bytes32 inputHash; // 32 bytes
        bytes32 outputHash; // 32 bytes
    }

    constructor() {
        owner = msg.sender;
    }

    modifier isOwner() {
        require(msg.sender == owner, "Caller is not owner");
        _;
    }

    modifier isPaused() {
        require(!paused, "Contract is paused");
        _;
    }

    event TaskCreated(
        bytes32 indexed uuid,
        address indexed submitter,
        string input,
        uint256 amount
    );

    event TaskProcessed(
        bytes32 indexed uuid,
        address indexed assignee,
        Status status,
        string output
    );

    /**
     * @dev Creates a new scraping task.
     * @param _uuid A unique identifier for the task, as bytes32.
     * @param _input The data for the task (e.g., a URL to scrape). The contract will store its hash.
     */
    function create(
        bytes32 _uuid,
        string memory _input
    ) external payable isPaused {
        require(_uuid != bytes32(0), "UUID cannot be empty");
        require(msg.value > 0, "Invalid amount");
        require(
            tasks[_uuid].submitter == address(0),
            "Task with this UUID already exists"
        );

        tasks[_uuid] = Task({
            submitter: payable(msg.sender),
            assignee: payable(address(0)),
            status: Status.Pending,
            amount: msg.value,
            inputHash: keccak256(bytes(_input)),
            outputHash: bytes32(0)
        });

        emit TaskCreated(_uuid, msg.sender, _input, msg.value);
    }

    /**
     * @dev Marks a task as complete or failed and pays the assignee.
     * @param _uuid The UUID of the task to complete.
     * @param _assignee The address of the node that performed the work.
     * @param _output The result of the scraping task. The contract will store its hash.
     * @param _success Whether the task was successful.
     */
    function complete(
        bytes32 _uuid,
        address payable _assignee,
        string memory _output,
        bool _success
    ) external isPaused {
        Task storage tsk = tasks[_uuid];
        require(tsk.submitter != address(0), "Invalid UUID");
        require(tsk.status == Status.Pending, "Task not pending");
        require(_assignee != address(0), "Assignee address cannot be zero");
        tsk.assignee = _assignee;

        if (_success) {
            tsk.status = Status.Completed;
            tsk.outputHash = keccak256(bytes(_output));
            if (tsk.amount > 0) {
                (bool sent, ) = _assignee.call{value: tsk.amount}("");
                require(sent, "Transfer failed");
            }
        } else {
            tsk.status = Status.Failed;
        }

        emit TaskProcessed(_uuid, tsk.assignee, tsk.status, _output);
    }

    /**
     * @dev Allows the owner to pause or unpause the contract.
     * @param _paused The desired pause state.
     */
    function setPaused(bool _paused) public isOwner {
        paused = _paused;
    }
}
