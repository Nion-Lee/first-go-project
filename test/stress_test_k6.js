import http from 'k6/http';
import { check, group, sleep } from 'k6';

export let options = {
    stages: [
        { duration: '30s', target: 250 },
        { duration: '90s', target: 600 },
        { duration: '1m', target: 500 },
        { duration: '2m', target: 80 },
        { duration: '30s', target: 0 },
    ],
    thresholds: {
        http_req_duration: ['p(99)<500', 'p(90)<250', 'p(50)<100'],
        http_req_failed: ['rate<0.1'], // 请求失败率低于10%
    },
};

function apiGet() {
    const url = 'http://localhost:8080/';
    let res = http.get(url);
    check(res, {
        'status is 200': (r) => r.status === 200,
        'content-type is application/json': (r) => r.headers['Content-Type'] === 'application/json',
        //'response body contains "Hello, World!"': (r) => r.body.indexOf('Hello, World!') !== -1,
    });
}

function apiPut() {
    const url = 'http://localhost:8080/Customer/Create';
    let payload = JSON.stringify(
    {     
        "UID": "", 
        "Name": "John Doe", 
        "Age": 30, 
        "Email": "johndoe@example.com" 
    });
    let params = { headers: { 'Content-Type': 'application/json' } };
    let res = http.put(url, payload, params);
    check(res, {
        'status is 200': (r) => r.status === 200,
        'content-type is application/json': (r) => r.headers['Content-Type'] === 'application/json',
        //'response body contains "Hello, World!"': (r) => r.body.indexOf('Hello, World!') !== -1,
    });
}

export default function () {
    group('API Group - GET', function () {
        apiGet();
    });

    group('API Group - PUT', function () {
        apiPut();
    });

    sleep(1);
}
