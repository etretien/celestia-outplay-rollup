


create-profile - создать профиль (имя, дата рождения, какой рукой играю, само-декларируемый рейтинг NTRP 1.0-6.0 через 0.5), изменение профиля пока не реализуем
create-challenge - вызов на игру со ставкой в токенах (может быть 0)
accept-challenge - принятие игры (блокирует ставку)
decline-challenge - отклонение
cancel-challenge - отмена уже принятой игры (и разблокировка ставок каждого участника)
submit-score - записать счет (= зафиксировать матч в блокчейне и распределить награду)
show-user-info - посмотреть инфо об участнике (включая данные профиля, рейтинг Elo и место в рейтинге, историю последних 5 матчей)
show-rating - вывести рейтинг игроков (отсортированный по Elo)



ignite scaffold message create-profile name dateOfBirth playingHand ntrpRating
ignite scaffold message create-challenge opponent stake
ignite scaffold message accept-challenge challengeId
ignite scaffold message decline-challenge challengeId
ignite scaffold message submit-score challengeId score

#queries
ignite scaffold message show-user-info userId
ignite scaffold message show-rating
ignite scaffold message my-challenges

ignite scaffold map profile owner name dateOfBirth playingHand ntrpRating elo --no-message
ignite scaffold map challenge challengeId stake challenger challenged status --no-message
ignite scaffold map match challengeId score winner loser --no-message


ignite scaffold query user-info user --response profile,matches

D259A6D596B9EEA62182CA2EF3B15CC8E91CC03E39044F36D5C23BC6A57E03D9

outplayd query tx --type=hash 7171E7A6F038B3AED34F86B8FB6842160A74ED52F724FA40C3AFEA7BE2BC17F9 --chain-id outplay --output json | jq -r '.raw_log'


outplayd tx outplay create-profile Dima 26.10.1988 right 3.5 --from=key-dima  --keyring-backend test
outplayd tx outplay create-profile Alex 26.10.1985 left 4.5 --from=key-alex  --keyring-backend test

outplayd tx outplay create-challenge outplay18s9wj4n8rejtewldnq5qmtv6uqgwp3axlxuf49 42069stake --from=key-alex --chain-id outplay --keyring-backend test
