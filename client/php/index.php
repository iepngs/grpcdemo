<?php

require_once __DIR__ . '/vendor/autoload.php';

use \Hello\HelloClient;
use \Hello\HelloReply;
use \Hello\HelloRequest;

// 创建客户端实例
$client = new HelloClient('127.0.0.1:50052', [
    'credentials' => Grpc\ChannelCredentials::createInsecure()
]);

// 组装请求参数
$request = new HelloRequest();
$request -> setName("iepngs");

// 发送请求
list($response, $status) = $client->SayHello($request)->wait();

if (0 != $status->code) {
    throw new \Exception($status->details, $status->code);
}

echo $response->getMessage() . PHP_EOL;