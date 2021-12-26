## This is a simple event receiver in cloudEvents format. To test it out, you can try the following curl command after you launched the application:

    curl -X POST \
        -H "content-type: application/json"  \
        -H "ce-specversion: 1.0"  \
        -H "ce-source: curl-command"  \
        -H "ce-type: curl.demo"  \
        -H "ce-id: 123-abc"  \
        -d '{"name":"Dave"}' localhost:8080