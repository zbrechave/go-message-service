package service

import (
	"errors"

	"github.com/zbrechave/go-mds/dao"

	"github.com/zbrechave/go-mds/entity"
)

const (
	MESSAGE_WAITING_CONFIRM = "待确认"
	MESSAGE_SENDING         = "发送中"
	PUBLIC_YES              = "是"
	PUBLIC_NO               = "否"
)

type messageService struct {
	dao *dao.MessageDao
}

/**
 * 预存储消息.
 */
func (messageService *messageService) saveMessageWaitingConfirm(message *entity.MessageEntity) (int64, error) {
	if message.GetMessageBody() == "" {
		return 0, errors.New("保存消息为空")
	}
	if message.GetConsumerQueue() == "" {
		return 0, errors.New("消息队列不能为空")
	}
	message.SetStatus(MESSAGE_WAITING_CONFIRM)
	message.SetAreadlyDead(PUBLIC_NO)
	message.SetMessageSendTimes(0)
	return messageService.dao.Insert(message)
}

/**
 * 确认并发送消息.
 */
func (messageService *messageService) confirmAndSendMessage(messageId string) error {
	return nil
}

/**
 * 存储并发送消息.
 */

func (messageService *messageService) saveAndSendMessage(message entity.MessageEntity) error {
	return nil
}

/**
 * 直接发送消息.
 */

func (messageService *messageService) directSendMessage(message entity.MessageEntity) error {
	return nil
}

/**
 * 重发消息.
 */

func (messageService *messageService) reSendMessage(message entity.MessageEntity) error {
	return nil
}

/**
 * 根据messageId重发某条消息.
 */

func (messageService *messageService) reSendMessageByMessageId(messageId string) error {
	return nil
}

/**
 * 将消息标记为死亡消息.
 */

func (messageService *messageService) setMessageToAreadlyDead(messageId string) error {
	return nil
}

/**
 * 根据消息ID获取消息
 */

func (messageService *messageService) getMessageByMessageId(messageId string) (entity.MessageEntity, error) {
	return entity.Message{}, nil
}

/**
 * 根据消息ID删除消息
 */

func (messageService *messageService) deleteMessageByMessageId(messageId string) error {
	return nil
}

/**
 * 重发某个消息队列中的全部已死亡的消息.
 */

func (messageService *messageService) reSendAllDeadMessageByQueueName(queueName string, batchSize int) error {
	return nil
}