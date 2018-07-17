from random import choice, random
from math import isclose


DEAL_AMOUNT_INTERVAL = (5, 10)  # interval for deal amount between one pair of traders
# possible trader actions
CHEAT = 0
COOPERATE = 1

TRADERS_AMOUNT = 60
# don't change the order
TRADERS_TYPES = (
    'altruist', 'trickster', 'dodger',
    'unpredictable', 'vindictive', 'quirky'
)


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


def correction_for_the_error(logic_func):
    def wrapper(*args):
        action = logic_func(*args)
        if isclose(random() - 0.05, 0, rel_tol=1e-6):
            return CHEAT if action else COOPERATE
        return action
    return wrapper


def get_logic_for_type(type):
    logics = (
        altruist_logic, trickster_logic, dodger_logic,
        unpredictable_logic, vindictive_logic, quirky_logic
    )

    type_logic_map = dict(zip(TRADERS_TYPES, logics))

    return correction_for_the_error(type_logic_map[type])


class Trader:
    def __init__(self, type):
        self.balance = 0
        self.type = type
        self._action_logic = get_logic_for_type(type)

    def __str__(self):
        return self.type

    def perform_action(self, opponent_turns):
        return self._action_logic(opponent_turns)


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


class Game:
    def __init__(self):
        self.turn = 0  # every turn is a year

        traders = []
        each_type_amount = TRADERS_AMOUNT // len(TRADERS_TYPES)
        for trader_type in TRADERS_TYPES:
            traders.extend([Trader(trader_type) for _ in range(each_type_amount)])

        self.traders = traders

    def make_turn(self):
        pass


if __name__ == '__main__':
    game_instance = Game()

    # debug
    for t in game_instance.traders:
        print(t)
