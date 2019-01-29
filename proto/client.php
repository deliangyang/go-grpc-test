<?php
/**
 * Created by ps.
 * User: phantom
 * Date: 2019/1/23 12:10
 * Email: 623601391@qq.com
 */

use Client\Hello;

require_once __DIR__ . '/vendor/autoload.php';

$client = new Hello('192.168.3.116:5000', [
    'credentials' => \Grpc\ChannelCredentials::createInsecure(),
]);

$request = new \Hello\HelloRequest();
$request->setName('1');

list($replay, $status) = $client->SayHello($request)->wait();
/**
 * @var $replay \Hello\HelloReply
 */
$getData = $replay->getMessage();
var_dump($getData);