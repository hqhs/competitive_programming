from collections import namedtuple


DEAL_AMOUNT_INTERVAL = (5, 10)
# possible trader actions
CHEAT = 0
COOPERATE = 1


def altruist_logic(*args):
    return COOPERATE


def trickster_logic(*args):
    # 'кидала'
    return CHEAT


def dodger_logic(opponent_turns=None):
    # 'хитрец'
    if not opponent_turns:
        return COOPERATE
    return opponent_turns[-1]


def unpredictable_logic(*args):
    # 'непредсказуемый'
    return choice(CHEAT, COOPERATE)


def vindictive_logic(opponent_turns=None):
    # 'злопамятный'
    # since COOPERATE is 1, bool(1) == True
    # all(opponent_turns) checks what opponent didn't chead
    # all([]) == True, so first turn is always cooperate
    return COOPERATE if all(opponent_turns) else CHEAT


def quirky_logic(opponent_turns=None):
    # 'ушлый'
    if len(opponent_turns) < 4:
        default = (COOPERATE, CHEAT, COOPERATE, COOPERATE)
        return default[len(opponent_turns)]
    else:
        decisive_moves = opponent_turns[:4]
        if all(decisive_moves):
            # play as dodger
            return dodger_logic(opponent_turns[4:])
        else:
            # play as trickster
            return trickster_logic()


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
