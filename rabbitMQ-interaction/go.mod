module rabbitmq

go 1.17

require github.com/rabbitmq/amqp091-go v1.2.0
require rabbit_utils v1.0.0
replace rabbit_utils v1.0.0 => ./rabbit_utils
