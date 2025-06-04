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
    }

    constructor() {
        owner = msg.sender;
    }

    modifier isTaskOwner(string memory _uuid) {
        Task memory tsk = tasks[_uuid];
        require(tsk.assignee == msg.sender, "Not submitters");
        require(tsk.status == Status.Claimed);
        _;
    }

    event TaskCreated(string uuid, string input);
    event TaskCompleted(string uuid, address assignee, string output);
    event TaskFailed(string uuid);

    function create(string memory _uuid, string memory _input) public payable {
        require(msg.value > 0, "Invalid amount");
        Task memory tsk = Task({
            uuid: _uuid,
            submitter: msg.sender,
            input: _input,
            output: "",
            status: Status.Pending,
            assignee: address(0),
            amount: msg.value
        });

        tasks[tsk.uuid] = tsk;
        emit TaskCreated(tsk.uuid, tsk.input);
    }

    function claim(string memory _uuid) external {
        require(
            tasks[_uuid].status == Status.Pending ||
                tasks[_uuid].status == Status.Failed,
            "Not pending or failed"
        );
        tasks[_uuid].status = Status.Claimed;
        tasks[_uuid].assignee = msg.sender;
    }

    function complete(
        string memory _uuid,
        string memory _output
    ) external payable isTaskOwner(_uuid) {
        tasks[_uuid].status = Status.Completed;
        tasks[_uuid].output = _output;
        emit TaskCompleted(_uuid, msg.sender, _output);
        (bool success, ) = tasks[_uuid].assignee.call{
            value: tasks[_uuid].amount
        }("");
        require(success == true, "Transfer failed");
    }

    function fail(string memory _uuid) external isTaskOwner(_uuid) {
        require(tasks[_uuid].status == Status.Claimed);
        require(tasks[_uuid].assignee != address(msg.sender));
        tasks[_uuid].status = Status.Failed;
        tasks[_uuid].assignee = address(0);
        emit TaskFailed(_uuid);
    }
}
