import { Bet, StyleObject } from '../types/index.d'
import BetCard from './BetCard'
import Subtitle from './Subtitle'

// OnGoingBets component. Displays site public bets.
function OnGoingBets() {
  // An array of bets
  const bets: Bet[] = [
    {
      id: 1,
      status: 'open',
      placer: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challenger: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
      moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 2,
      status: 'open',
      placer: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challenger: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
      moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 3,
      status: 'open',
      placer: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challenger: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
      moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 4,
      status: 'open',
      placer: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challenger: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
      moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 5,
      status: 'open',
      placer: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challenger: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
      moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
    {
      id: 6,
      status: 'open',
      placer: '0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC',
      challenger: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
      moderator: '0x39249126d90671284cd06495d19C04DD0e54d371',
      description: 'In 2022 there will be 2000 electric cars accidents',
      terms: 'Has to be in the us.',
      expirationDate: 'Fri Sep 16 2022',
      amount: 30,
    },
  ]

  // Centralized all UI styles in one place for improve in readability.
  const styles: StyleObject = {
    publicBets: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'start',
      alignItems: 'flex-start',
      gap: '24px',
      padding: '32px 62px',
    },
  }
  return (
    <>
      <Subtitle showSearch text="Site Ongoing Bets" />
      <section style={styles.publicBets}>
        {bets.map((bet) => (
          <BetCard key={bet.id} isDetail={false} bet={bet} />
        ))}
      </section>
    </>
  )
}

export default OnGoingBets
