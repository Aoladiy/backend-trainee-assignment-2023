All tasks complete
<ol>
    <li>main task</li>
    <li>first additional task</li>
    <li>second additional task</li>
    <li>third additional task</li>
</ol>
How to start project
<ol>
    <li>create file .env with content like in .env.example</li>
    <li>(optional) change ports in docker-compose.yml if they are already in use</li>
    <li>docker-compose build</li>
    <li>docker-compose up -d</li>
    <li>
        test the app
        <ul>
            <li>
                if you have "Could not get response" problem when trying<br>
                to send http request, just wait a few moments<br>
                (something about 10-60 seconds).<br>
                It may take a time to start db
            </li>
            <li>
                you can import<br>
                examples of http requests into postman<br>
                using backend-trainee-assignment-2023.postman_collection.json<br>
                in the project
            </li>
            <li>
                pay attention to the http body in several requests<br>
                there are some parameters or/and response
            </li>
            <li>
                note that if you want to download csv file from <br>
                get user log request you should tap Send and Download<br>
                instead of just Send
            </li>
        </ul>
    </li>
</ol>
