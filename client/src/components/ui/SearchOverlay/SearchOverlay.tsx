import React from "react";

import styles from "./SearchOverlay.module.css";
import SearchInput from "./SearchInput";



const SearchOverlay = () => {
  return (
    <div>
      <div className={styles['justify-center']}>
        <h1 className={styles.heading}>
          Friends<span className={styles.hk}>HK</span>
        </h1>
        <h2>Hledat</h2>
    <SearchInput placeholder="Akce, festivaly..."  />
    <div className={styles.mediumGap}></div>
        <h2>Pozvánky</h2>

      </div>
    </div>
  );
};

export default SearchOverlay;

