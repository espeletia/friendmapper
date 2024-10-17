import { useState } from "react";
import { Books, Castle, LuckyClover, Palete, Song, StrongMan } from "../../svg";
import Selector from "./components/Selector";
import styles from "./Filters.module.css";

const Filters = () => {
  const [selectors, setSelectors] = useState([
    { selected: true, name: "Vše", icon: <Palete />, deseletsAll: true },
    { selected: false, name: "Festivaly", icon: <Castle /> },
    { selected: false, name: "Divadla", icon: <Song /> },
    { selected: false, name: "Kina", icon: <StrongMan /> },
    { selected: false, name: "Sport", icon: <Books /> },
    { selected: false, name: "Lucky", icon: <LuckyClover /> },
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
              {selector.name}
            </>
          </Selector>
        ))}
      </div>
    </div>
  );
};

export default Filters;
