<?php
/**
 * Created by ps.
 * User: phantom
 * Date: 2019/1/23 12:05
 * Email: 623601391@qq.com
 */
namespace Client;

use Grpc\BaseStub;
use Hello\HelloRequest;

class Hello extends BaseStub
{
    public function __construct(string $hostname, array $opts, $channel = null)
    {
        parent::__construct($hostname, $opts, $channel);
    }

    public function SayHello(HelloRequest $argument, $metadata = [], $options = [])
    {
        return $this->_simpleRequest('/hello.Greeter/SayHello',
            $argument,
            ['\Hello\HelloReply', 'decode'],
            $metadata,
            $options
        );
    }
}