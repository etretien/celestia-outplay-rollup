# Outplay Tennis Blockchain App

Outplay is the first application for amateur athletes that allows you to receive cash rewards for playing the sport you love. You can also find partners and track your progress in your favorite sport within the app. Build your own sports community and receive extra rewards for having around you the people who share the same passion as you.

This repository contains the blockchain backend for the Outplay Tennis app, built using Cosmos SDK and Rollkit. The app allows users to create profiles, challenge other users to matches, accept or decline challenges, and submit match scores.

## Features

- **Create Profile**: Users can create a profile with their name, date of birth, playing hand, and NTRP rating. The profile also includes an Elo rating, which is initially calculated based on the NTRP rating. The command to create a profile is `create-profile`.

- **Create Challenge**: Users can challenge other users to a match by specifying the stake (in tokens) for the match. The command to create a challenge is `create-challenge`.

- **Accept Challenge**: The challenged user can accept the challenge. Upon acceptance, the same amount of tokens as the stake is locked from their account. The command to accept a challenge is `accept-challenge`.

- **Decline Challenge**: The challenged user can also decline the challenge. If a challenge is declined, the challenger's locked tokens are unlocked. The command to decline a challenge is `decline-challenge`.

- **Submit Score**: After a match, the winner can submit the score. The score is verified and saved, and the total stake (from both users) is transferred to the winner. The Elo ratings of both users are also updated based on the match result. The command to submit a score is `submit-score`.

## Building and Running

To build the application, you can use the `./init.sh` command from the root directory of the repository.

You need a Celestia light node running locally with some TIA balance on it.

## Examples

Here are some examples of how to use the commands using the Outplay CLI:

```bash
# Create profiles
outplayd tx outplay create-profile Dima 26.10.1988 right 3.5 --from=dima --keyring-backend test
outplayd tx outplay create-profile Alex 26.10.1985 left 4.5 --from=alex --keyring-backend test

# Create a challenge
outplayd tx outplay create-challenge outplay1a7c5rmu7rzdwal7kwcmm05txxm0fq30d8rz8ju 80000 --from=alex --keyring-backend test

# Decline a challenge
outplayd tx outplay decline-challenge 3f5ee6e9310e502bfb623ff493e651c6ca9ac691c36740f0552b2cdbd5ff9791 --from=dima --keyring-backend test

# Accept a challenge
outplayd tx outplay accept-challenge 8d8b0a01bc681979f5ad80a1e4516758d7db5cf93e61b6f3bc15d5e4b30bae9f --from=dima --keyring-backend test

# Submit a score
outplayd tx outplay submit-score 8d8b0a01bc681979f5ad80a1e4516758d7db5cf93e61b6f3bc15d5e4b30bae9f "1:6,7:6,1:6" --from=dima --keyring-backend test

# List user profiles
outplayd query outplay list-profile 

Output example (note the update Elo ratings):
profile:
- dateOfBirth: 26.10.1985
  elo: "2387"
  index: bd10d5a7012a1cd2d9e8d24cd60527fba3ef1e7698cfd08a912a7c3ebe0c18c3
  name: Alex
  ntrpRating: "4.5"
  owner: outplay1gnd76vf2sezajrm98mj2k7l2wa5y9frzfezg45
  playingHand: left
- dateOfBirth: 26.10.1988
  elo: "2029"
  index: fa3867f5aca1185974ee975d60983e4e8be834691bc82e65bb1a9ef3ef698f9a
  name: Dima
  ntrpRating: "3.5"
  owner: outplay1ws8m4ch8m350xz6dck4gvh9py9za6qv6hd504p
  playingHand: right

# List challenges
outplayd query outplay list-challenges

Output example:
challenge:
- challengeId: 5896a89a815199268202d8c85462409fd262cff96b71e38eff72fcfe0e4a5449
  challenged: outplay1ws8m4ch8m350xz6dck4gvh9py9za6qv6hd504p
  challenger: outplay1gnd76vf2sezajrm98mj2k7l2wa5y9frzfezg45
  index: 5896a89a815199268202d8c85462409fd262cff96b71e38eff72fcfe0e4a5449
  stake: "80000.00"
  status: accepted

# List matches:
outplayd query outplay list-matches

Output example (winner and loser are determined automatically):
match:
- challengeId: ""
  index: 5896a89a815199268202d8c85462409fd262cff96b71e38eff72fcfe0e4a5449
  loser: outplay1gnd76vf2sezajrm98mj2k7l2wa5y9frzfezg45
  score: 1:6,7:6,1:6
  winner: outplay1ws8m4ch8m350xz6dck4gvh9py9za6qv6hd504p
```

## Future Work
This is a prototype and there are many potential improvements and additional features that could be added. Some possibilities include:

- Adding more detailed match information, such as the date and location of the match.
- Allowing users to submit evidence of the match result, such as photos or videos.
- Implementing a dispute resolution mechanism for contested match results.
- Adding support for different types of matches (e.g., singles, doubles) and different scoring systems.
- Integrating with a front-end app for a better user experience (for now only CLI has been implemented).

We welcome contributions and feedback. Please feel free to open an issue or submit a pull request if you have any suggestions or find any bugs.
