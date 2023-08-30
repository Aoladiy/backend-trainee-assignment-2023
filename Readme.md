<ol>
    <li>create file .env with content like in .env.example</li>
    <li>(optional) change ports in docker-compose.yml if they are already in use</li>
    <li>docker-compose build</li>
    <li>docker-compose up -d</li>
    <li>try to start container with app when container with db will be ready
        <ul>
            <li>docker ps -a (with this command you can see container ids)</li>
            <li>docker start {insert here container id which have image backend-trainee-assignment-2023-go}</li>
        </ul>
    </li>
    <li>when container with db will be ready
        <ul>
            <li>docker ps -a (with this command you can see container ids)</li>
            <li>docker exec -it {insert here container id which have image mysql:8.0} mysql -uroot -proot</li>
            <li>
<pre>
use dockermysql;
CREATE EVENT cleanup_expired_segments
ON SCHEDULE EVERY 1 SECOND
ON COMPLETION PRESERVE
DO DELETE FROM segments_users WHERE expiration_time <= NOW();
</pre>
            </li>
            <li>
<pre>
CREATE TRIGGER segments_users_insert_trigger
AFTER INSERT ON segments_users
FOR EACH ROW INSERT INTO segments_users_log (user_id, segment_id, action, datetime) VALUES (NEW.user_id, NEW.segment_id, 'insert', NOW());
</pre>
            </li>
            <li>
<pre>
CREATE TRIGGER segments_users_delete_trigger
AFTER DELETE ON segments_users
FOR EACH ROW INSERT INTO segments_users_log (user_id, segment_id, action, datetime) VALUES (OLD.user_id, OLD.segment_id, 'delete', NOW());
</pre>
            </li>
            <li>
                exit
            </li>
        </ul>
    </li>
</ol>