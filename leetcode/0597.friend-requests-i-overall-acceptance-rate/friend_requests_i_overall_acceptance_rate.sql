Create table If Not Exists friend_request
(
    sender_id    INT  NOT NULL,
    send_to_id   INT  NULL,
    request_date DATE NULL
);
Create table If Not Exists request_accepted
(
    requester_id INT  NOT NULL,
    accepter_id  INT  NULL,
    accept_date  DATE NULL
);
Truncate table friend_request;
insert into friend_request (sender_id, send_to_id, request_date)
values ('1', '2', '2016/06/01');
insert into friend_request (sender_id, send_to_id, request_date)
values ('1', '3', '2016/06/01');
insert into friend_request (sender_id, send_to_id, request_date)
values ('1', '4', '2016/06/01');
insert into friend_request (sender_id, send_to_id, request_date)
values ('2', '3', '2016/06/02');
insert into friend_request (sender_id, send_to_id, request_date)
values ('3', '4', '2016/06/09');
Truncate table request_accepted;
insert into request_accepted (requester_id, accepter_id, accept_date)
values ('1', '2', '2016/06/03');
insert into request_accepted (requester_id, accepter_id, accept_date)
values ('1', '3', '2016/06/08');
insert into request_accepted (requester_id, accepter_id, accept_date)
values ('2', '3', '2016/06/08');
insert into request_accepted (requester_id, accepter_id, accept_date)
values ('3', '4', '2016/06/09');
insert into request_accepted (requester_id, accepter_id, accept_date)
values ('3', '4', '2016/06/10');

# Write your MySQL query statement below
SELECT ROUND(IFNULL(
                         (SELECT COUNT(*) FROM (SELECT DISTINCT requester_id, accepter_id FROM request_accepted) AS A)
                         /
                         (SELECT COUNT(*) FROM (SELECT DISTINCT sender_id, send_to_id FROM friend_request) AS B), 0),
             2) AS accept_rate;
