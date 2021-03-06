// Copyright 2020 The klaytn Authors
// This file is part of the klaytn library.
//
// The klaytn library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The klaytn library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the klaytn library. If not, see <http://www.gnu.org/licenses/>.

package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

const (
	topicProjectName = "klaytn"
	topicServiceName = "chaindatafetcher"
	topicVersion     = "v1"
)

const (
	DefaultReplicas             = 1
	DefaultPartitions           = 1
	DefaultTopicEnvironmentName = "local"
	DefaultTopicResourceName    = "en-0"
)

type KafkaConfig struct {
	SaramaConfig         *sarama.Config // kafka client configurations.
	Brokers              []string       // Brokers is a list of broker URLs.
	TopicEnvironmentName string
	TopicResourceName    string
	Partitions           int32 // Partitions is the number of partitions of a topic.
	Replicas             int16 // Replicas is a replication factor of kafka settings. This is the number of the replicated partitions in the kafka cluster.
}

func GetDefaultKafkaConfig() *KafkaConfig {
	// TODO-ChainDataFetcher add more configuration if necessary
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Version = sarama.MaxVersion
	return &KafkaConfig{
		SaramaConfig:         config,
		TopicEnvironmentName: DefaultTopicEnvironmentName,
		TopicResourceName:    DefaultTopicResourceName,
		Partitions:           DefaultPartitions,
		Replicas:             DefaultReplicas,
	}
}

func (c *KafkaConfig) getTopicName(event string) string {
	return fmt.Sprintf("%v.%v.%v.%v.%v.%v", c.TopicEnvironmentName, topicProjectName, topicServiceName, c.TopicResourceName, event, topicVersion)
}
