package mgr

import (
	"awesomeProject/src/Proto"
	"awesomeProject/src/entry/interface"
	"sync"
)

var pActionsMgr *actionsMgr
var onceAction sync.Once

func GetActionsMgr() *actionsMgr {
	onceAction.Do(func() {
		pActionsMgr = & actionsMgr{}
		pActionsMgr.actionMap = make(map[uint32] *_interface.IAction)
		pActionsMgr.ruleLists = make(map[string] []*_interface.IRule)
	})
	return pActionsMgr
}

type actionsMgr struct {
	actions []*_interface.IAction //按执行顺序存放 ，没有规则是按顺序执行
	actionMap map[uint32]*_interface.IAction
	ruleLists map[string][]*_interface.IRule
	curRuleList string
}
// 复制一个新action
func (pActionsMgr *actionsMgr) GetFirstAction() *_interface.IAction {
	return pActionsMgr.actions[0]
}

func (pActionsMgr *actionsMgr) GetNetAction(curAction _interface.IAction) *_interface.IAction {
	var rule _interface.IRule = nil
	if len(pActionsMgr.GetCurRuleList()) != 0 {
		// 尝试寻找rule
		for _,k := range pActionsMgr.GetCurRuleList() {
			pRule := *k
			if pRule.GetCurActionId() == curAction.GetActionState() || pRule.GetFinishCode() == curAction.GetActionState() {
				rule = pRule
				break
			}
		}
	}
	// 没有适应的跳转规则，查找下一个
	if rule == nil {
		iAction := pActionsMgr.actions[curAction.GetIndex()]
		return iAction
	}
	return pActionsMgr.actionMap[rule.GetNextActionId()]
}

func (pActionsMgr *actionsMgr) AddAction(iAction *_interface.IAction) {
	if iAction == nil{
		return
	}
	(*iAction).SetActionState(Proto.ACTION_STATE_NONE)
	(*iAction).SetIndex(uint8(len(pActionsMgr.actions)))
	pActionsMgr.actions = append(pActionsMgr.actions, iAction)
	pActionsMgr.actionMap[(*iAction).GetActionId()] = iAction
}

func (pActionsMgr *actionsMgr) ConfigAction(actionName uint32, attributes []string) {

}

func (pActionsMgr *actionsMgr) GetAllActions() []*_interface.IAction {
	return pActionsMgr.actions
}

func (pActionsMgr *actionsMgr) SetCurRuleList(ruleListName string) {
	pActionsMgr.curRuleList = ruleListName
}

func (pActionsMgr *actionsMgr) AddRule(rule *_interface.IRule) {
	ruleList,ok := pActionsMgr.ruleLists[pActionsMgr.curRuleList]
	// ruleList 并不是指针
	if ok {
		pActionsMgr.ruleLists[pActionsMgr.curRuleList] = append(ruleList, rule)
	} else {
		ruleList = make([]*_interface.IRule, 1)
		pActionsMgr.ruleLists[pActionsMgr.curRuleList] = append(ruleList, rule)
	}
}

func (pActionsMgr *actionsMgr) GetCurRuleList() []*_interface.IRule {
	return pActionsMgr.ruleLists[pActionsMgr.curRuleList]
}