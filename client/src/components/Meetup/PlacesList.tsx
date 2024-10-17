import React from 'react';
import styles from './PlacesList.module.css';

const places = [
  { icon: '🎨', name: 'Divadlo', description: 'Festival', checked: false },
  { icon: '🎵', name: 'Divadlo', description: 'Festival', checked: true },
  { icon: '🎵', name: 'Divadlo', description: 'Festival', checked: false },
  { icon: '🏋️', name: 'Divadlo', description: 'Festival', checked: false },
  { icon: '🏰', name: 'Divadlo', description: 'Festival', checked: false },
];

const PlacesList = () => {
  return (
    <div className={styles.container}>
      <h2 className={styles.title}>Místa pro vás</h2>
      <ul className={styles.list}>
        {places.map((place, index) => (
          <li key={index} className={styles.item}>
            <span className={styles.icon}>{place.icon}</span>
            <div className={styles.info}>
              <span className={styles.name}>{place.name}</span>
              <span className={styles.description}>{place.description}</span>
            </div>
            <div className={styles.checkbox}>
              {place.checked ? '✔️' : '⬜️'}
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default PlacesList;