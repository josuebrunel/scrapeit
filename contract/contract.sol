// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

contract Scraper {
    address public owner;
    enum Status {
        Pending,
        Claimed,
        Completed,
        Failed
    }

    mapping(string => Task) public tasks;

    struct Task {
        string uuid;
        address submitter;
        string input;
        string output;
        Status status;
        address assignee;
        uint256 amount;
        uint256 created_at;
        uint256 updated_at;
    }

    constructor() {
        owner = msg.sender;
    }

    modifier isTaskOwner(string memory _uuid) {
        Task memory tsk = tasks[_uuid];
        require(
            tasks[_uuid].assignee == msg.sender,
            "Caller is not task assignee"
        );
        require(
            tasks[_uuid].status == Status.Claimed,
            "Task not in Claimed state"
        );
        _;
    }

    event TaskCreated(
        string uuid,
        address indexed submitter,
        string input,
        uint256 amount
    );
    event TaskClaimed(string uuid, address indexed assignee);
    event TaskCompleted(string uuid, address indexed assignee, string output);
    event TaskFailed(string uuid, address indexed assigneeWhoFailed); // Task returned to pool by assignee

    function create(string memory _uuid, string memory _input) public payable {
        require(bytes(_uuid).length > 0, "UUID cannot be empty");
        require(msg.value > 0, "Invalid amount");
        require(
            tasks[_uuid].submitter == address(0),
            "Task with this UUID already exists"
        );

        Task memory tsk = Task({
            uuid: _uuid,
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
        emit TaskCreated(tsk.uuid, msg.sender, tsk.input, msg.value);
    }

    function claim(string memory _uuid) external {
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
    ) external isTaskOwner(_uuid) {
        require(bytes(_uuid).length > 0, "UUID cannot be empty");
        Task storage tsk = tasks[_uuid];

        tsk.status = Status.Completed;
        tsk.output = _output;
        tsk.updated_at = block.timestamp;
        emit TaskCompleted(_uuid, msg.sender, _output);

        // Transfer payment to the assignee
        (bool success, ) = tsk.assignee.call{value: tsk.amount}("");
        require(success, "Transfer failed");
    }

    function fail(string memory _uuid) external isTaskOwner(_uuid) {
        require(bytes(_uuid).length > 0, "UUID cannot be empty");
        Task storage tsk = tasks[_uuid];

        tsk.status = Status.Failed;
        tsk.assignee = address(0); // Reset assignee so task can be claimed again
        tsk.updated_at = block.timestamp;
        emit TaskFailed(_uuid, msg.sender);
    }
}
