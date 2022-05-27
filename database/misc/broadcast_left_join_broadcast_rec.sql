-- Join tables to get broadcast with N number of receipients
-- SELECT broadcast_id, type, title, content, creation_date, deadline, creator, 
-- broadcast_recipients_id, related_broadcast, recipient, acknowledged, rejected
SELECT *
FROM operations_ecosystem.broadcast 
LEFT JOIN operations_ecosystem.broadcast_recepients
ON broadcast_id = related_broadcast
WHERE broadcast_id IN (
SELECT broadcast_id
FROM operations_ecosystem.broadcast 
LEFT JOIN operations_ecosystem.broadcast_recepients 
ON broadcast_id = related_broadcast
WHERE  acknowledged=false
GROUP BY  broadcast_id
HAVING COUNT(broadcast_id) > 2
);
