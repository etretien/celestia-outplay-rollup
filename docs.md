Please add tech details fo README (add more describing text, all in English), given raw details below. 

Messages:
create-profile - создать профиль (имя, дата рождения, какой рукой играю, само-декларируемый рейтинг NTRP 1.0-6.0 через 0.5), изменение профиля пока не реализуем
create-challenge - вызов на игру со ставкой в токенах (может быть 0)
accept-challenge - принятие игры (блокирует ставку)
decline-challenge - отклонение
submit-score - записать счет (= зафиксировать матч в блокчейне и распределить награду)

ignite scaffold message create-profile name dateOfBirth playingHand ntrpRating
ignite scaffold message create-challenge opponent stake
ignite scaffold message accept-challenge challengeId
ignite scaffold message decline-challenge challengeId
ignite scaffold message submit-score challengeId score

Maps:
ignite scaffold map profile owner name dateOfBirth playingHand ntrpRating elo --no-message
ignite scaffold map challenge challengeId stake challenger challenged status --no-message
ignite scaffold map match challengeId score winner loser --no-message


Examples:
outplayd tx outplay create-profile Dima 26.10.1988 right 3.5 --from=dima --keyring-backend test
outplayd tx outplay create-profile Alex 26.10.1985 left 4.5 --from=alex --keyring-backend test

outplayd tx outplay create-challenge outplay1a7c5rmu7rzdwal7kwcmm05txxm0fq30d8rz8ju 80000 --from=alex --keyring-backend test
outplayd tx outplay decline-challenge 3f5ee6e9310e502bfb623ff493e651c6ca9ac691c36740f0552b2cdbd5ff9791 --from=dima --keyring-backend test
outplayd tx outplay accept-challenge 8d8b0a01bc681979f5ad80a1e4516758d7db5cf93e61b6f3bc15d5e4b30bae9f --from=dima --keyring-backend test
outplayd tx outplay submit-score 8d8b0a01bc681979f5ad80a1e4516758d7db5cf93e61b6f3bc15d5e4b30bae9f "1:6,7:6,1:6" --from=dima --keyring-backend test

Example queries:
List user profiles:
outplayd query outplay list-profile   
output:
profile:
- dateOfBirth: 26.10.1985
  elo: "2390"
  index: 59a51c33fa2fab5ea1ba20ab251b3b94e56e84d74b9c78b8bf36d219357eb846
  name: Alex
  ntrpRating: "4.5"
  owner: outplay1lqfy6uaefp9emv4aw0jfhl6ed77n0872n5jym0
  playingHand: left
- dateOfBirth: 26.10.1988
  elo: "2046"
  index: b395995e9b3c1f2f37698b32db041673053345709afcbb36f86dc63819671efa
  name: Dima
  ntrpRating: "3.5"
  owner: outplay1a7c5rmu7rzdwal7kwcmm05txxm0fq30d8rz8ju
  playingHand: right

List challenges:
outplayd query outplay list-challenges
output:
challenge:
- challengeId: 8d8b0a01bc681979f5ad80a1e4516758d7db5cf93e61b6f3bc15d5e4b30bae9f
  challenged: outplay1a7c5rmu7rzdwal7kwcmm05txxm0fq30d8rz8ju
  challenger: outplay1lqfy6uaefp9emv4aw0jfhl6ed77n0872n5jym0
  index: 8d8b0a01bc681979f5ad80a1e4516758d7db5cf93e61b6f3bc15d5e4b30bae9f
  stake: "80000.00"
  status: finished

List matches:
outplayd query outplay list-matches
match:
- challengeId: ""
  index: 8d8b0a01bc681979f5ad80a1e4516758d7db5cf93e61b6f3bc15d5e4b30bae9f
  loser: outplay1lqfy6uaefp9emv4aw0jfhl6ed77n0872n5jym0
  score: 1:6,7:6,1:6
  winner: outplay1a7c5rmu7rzdwal7kwcmm05txxm0fq30d8rz8ju
- challengeId: ""
  index: 930a825025fa26d78fb900aff51a019c42f78b69cdaa43c4c194027b802d7061
  loser: outplay1a7c5rmu7rzdwal7kwcmm05txxm0fq30d8rz8ju
  score: 6:4,6:4
  winner: outplay1lqfy6uaefp9emv4aw0jfhl6ed77n0872n5jym0
