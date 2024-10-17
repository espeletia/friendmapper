import { useState } from "react";
import { Castle, Palete, Song, StrongMan } from "../../svg";
import Selector from "./components/Selector";
import styles from "./Filters.module.css";

const Filters = () => {
  const [selectors, setSelectors] = useState([
    { selected: true, name: "Vše", icon: null, deseletsAll: true },
    { selected: false, name: "Festivaly", icon: <Song /> },
    { selected: false, name: "Divadla", icon: <Palete /> },
    { selected: false, name: "Kina", icon: <Castle />},
    { selected: false, name: "Sport", icon: <StrongMan />  },
  ]);

  const handleChange = (index: number) => {
    setSelectors((prevSelectors) => {
      const isDeselectAllSelected = prevSelectors[index]?.deseletsAll;

      return prevSelectors.map((selector, i) => {
        if (isDeselectAllSelected) {
          return {
            ...selector,
            selected: i === index,
          };
        } else {
          return {
            ...selector,
            selected:
              i === index
                ? !selector.selected
                : selector.deseletsAll
                  ? false
                  : selector.selected,
          };
        }
      });
    });
  };

  return (
    <div className={styles.container}>
      <div className={styles.filterContainer}>
        {selectors.map((selector, index) => (
          <Selector
            key={index}
            index={index}
            selected={selector.selected}
            handleChange={handleChange}
          >
            <>
              {selector.icon}
              {selector.icon && <span className={styles.gap}></span>}
              {selector.name}
            </>
          </Selector>
        ))}
      </div>
    </div>
  );
};

export default Filters;
