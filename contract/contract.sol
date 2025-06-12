// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

contract Scraper {
    address public owner;
    bool public paused;
    enum Status {
        Pending,
        Completed,
        Failed
    }

    mapping(string => Task) public tasks;

    struct Task {
        string uuid;
        string user_id;
        Status status;
        address payable assignee;
        address payable submitter;
        string input;
        string output;
        uint256 amount;
        uint256 created_at;
        uint256 updated_at;
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
        string uuid,
        string user_id,
        address indexed submitter,
        string input,
        uint256 amount
    );
    event TaskProcessed(
        string uuid,
        address indexed assignee,
        Status status,
        string output
    );

    function create(
        string memory _uuid,
        string memory _user_id,
        string memory _input
    ) public payable isPaused {
        require(bytes(_uuid).length > 0, "UUID cannot be empty");
        require(msg.value > 0, "Invalid amount");
        require(
            tasks[_uuid].submitter == address(0),
            "Task with this UUID already exists"
        );

        Task memory tsk = Task({
            uuid: _uuid,
            user_id: _user_id,
            submitter: payable(msg.sender),
            input: _input,
            output: "",
            status: Status.Pending,
            assignee: payable(address(0)),
            amount: msg.value,
            created_at: block.timestamp,
            updated_at: block.timestamp
        });

        tasks[tsk.uuid] = tsk;
        emit TaskCreated(
            tsk.uuid,
            tsk.user_id,
            tsk.submitter,
            tsk.input,
            msg.value
        );
    }

    function complete(
        string memory _uuid,
        address payable _assignee,
        string memory _output,
        bool _success
    ) external isPaused {
        require(bytes(_uuid).length > 0, "UUID cannot be empty");
        Task storage tsk = tasks[_uuid];
        require(tsk.status != Status.Completed, "Task already completed");
        require(_assignee != address(0), "Assignee address cannot be zero");
        tsk.assignee = _assignee;

        if (!_success) {
            tsk.status = Status.Failed;
        } else {
            tsk.status = Status.Completed;
            tsk.output = _output;
        }
        tsk.updated_at = block.timestamp;

        if (_success && tsk.amount > 0) {
            (bool success, ) = tsk.assignee.call{value: tsk.amount}("");
            require(success, "Transfer failed");
        }

        emit TaskProcessed(_uuid, tsk.assignee, tsk.status, _output);
    }

    function setPaused(bool _paused) public isOwner {
        paused = _paused;
    }
}
