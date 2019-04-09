package impl

type Rule struct {
	curAction uint32 // 直接从cur 切换到next
	finishCode uint32 // 当满足结束code时 切换
	nextAction uint32
}

func (rule *Rule) GetNextActionId() uint32 {
	return rule.nextAction
}

func (rule *Rule) GetCurActionId() uint32 {
	return rule.curAction
}

func (rule *Rule) GetFinishCode() uint32 {
	return rule.finishCode
}

func (rule *Rule) SetCurAction(curAction uint32) {
	rule.curAction = curAction
}

func (rule *Rule) SetFinishCode(finishCode uint32) {
	rule.finishCode = finishCode
}

func (rule *Rule) SetNextAction(nextAction uint32) {
	rule.nextAction = nextAction
}
