// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

contract Scraper {
    address public owner;
    bool public paused;
    enum Status {
        Pending,
        Claimed,
        Completed,
        Failed
    }

    mapping(string => Task) public tasks;

    struct Task {
        string uuid;
        string user_id;
        Status status;
        address assignee;
        address submitter;
        string input;
        string output;
        uint256 amount;
        uint256 created_at;
        uint256 updated_at;
    }

    constructor() {
        owner = msg.sender;
    }

    modifier isTaskClaimed(string memory _uuid) {
        Task memory tsk = tasks[_uuid];
        require(
            tasks[_uuid].status != Status.Completed,
            "Task not in Claimed state"
        );
        _;
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
    event TaskClaimed(string uuid, address indexed assignee);
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
        require(tasks[_uuid].status == Status.Pending);

        Task memory tsk = Task({
            uuid: _uuid,
            user_id: _user_id,
            submitter: msg.sender,
            input: _input,
            output: "",
            status: Status.Pending,
            assignee: address(0),
            amount: msg.value,
            created_at: block.timestamp,
            updated_at: block.timestamp
        });

        tasks[tsk.uuid] = tsk;
        emit TaskCreated(
            tsk.uuid,
            tsk.user_id,
            msg.sender,
            tsk.input,
            msg.value
        );
    }

    function claim(string memory _uuid) external isPaused {
        require(bytes(_uuid).length > 0, "UUID cannot be empty");
        Task storage tsk = tasks[_uuid];
        require(
            tsk.status == Status.Pending || tsk.status == Status.Failed,
            "Not pending or failed"
        );
        require(
            tsk.submitter != msg.sender,
            "Submitter cannot claim their own task"
        );

        tsk.status = Status.Claimed;
        tsk.assignee = msg.sender;
        tsk.updated_at = block.timestamp;
        emit TaskClaimed(_uuid, msg.sender);
    }

    function complete(
        string memory _uuid,
        string memory _output
    ) external isTaskClaimed(_uuid) isPaused {
        require(bytes(_uuid).length > 0, "UUID cannot be empty");
        Task storage tsk = tasks[_uuid];

        tsk.status = Status.Completed;
        tsk.output = _output;
        tsk.updated_at = block.timestamp;
        emit TaskProcessed(_uuid, msg.sender, Status.Completed, _output);

        // Transfer payment to the assignee
        (bool success, ) = tsk.assignee.call{value: tsk.amount}("");
        require(success, "Transfer failed");
    }

    function fail(string memory _uuid) external isTaskClaimed(_uuid) isPaused {
        require(bytes(_uuid).length > 0, "UUID cannot be empty");
        Task storage tsk = tasks[_uuid];

        tsk.status = Status.Failed;
        tsk.assignee = address(0); // Reset assignee so task can be claimed again
        tsk.updated_at = block.timestamp;
        emit TaskProcessed(_uuid, msg.sender, Status.Failed, "");
    }

    function setPaused(bool _paused) public isOwner {
        paused = _paused;
    }
}
