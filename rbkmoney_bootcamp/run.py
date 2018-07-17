from collections import namedtuple


DEAL_AMOUNT_INTERVAL = (5, 10)
# possible trader actions
CHEAT = 0
COOPERATE = 1


class Trader:
    def __init__(self):
        self.balance = 0

    def perform_action(self):
        pass


def make_deal(trader1, trader2):
    trader1_action, trader2_action = trader1.perform_action(), trader2.perform_action()

    if trader1_action == CHEAT and trader2_action == CHEAT:
        trader1.balance += 2
        trader2.balance += 2
    elif trader1_action == COOPERATE and trader2_action == COOPERATE:
        trader1.balance += 4
        trader2.balance += 4
    elif trader1_action == COOPERATE and trader2_action == CHEAT:
        trader1.balance += 1
        trader2.balance += 5
    elif trader1_action == CHEAT and trader2_action == COOPERATE:
        trader1.balance += 5
        trader2.balance += 1
