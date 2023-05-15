


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
