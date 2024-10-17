import { useNavigate } from "react-router-dom";
import { Castle, CrossBlack, Palete, Song, StrongMan } from "../../svg";
import CustomButton from "../ui/AuthDialog/CustomButton";

import styles from "./Meetup.module.css";
import InputField from "../ui/AuthDialog/InputField";
import SelectField from "../ui/AuthDialog/SelectField";
import { useState } from "react";
import Selector from "./Selector";
import Filters from "../Filters/Filters";
import cx from "classnames";
import PlacesList from "./PlacesList";

interface Props {
  closeCallback: VoidFunction;
}

const Meetup = (props: Props) => {
    const [phase, setPhase] = useState(0);

    const [invited1, setInvited1] = useState(false);
    const [invited2, setInvited2] = useState(false);
    

  
  return (
    <div className={cx(styles.container, {[styles.container3]: phase === 1, [styles.container4]: phase === 2, [styles.container5]: phase === 3})}>
      <div className={styles.header}>
        <h2>Nový meetup</h2>
        <button onClick={props.closeCallback}>
          <CrossBlack />
        </button>
      </div>
      {
        phase === 0 && ( <><InputField type="text" title="Jméno meetupu" placeholder="Zadejte jméno..."></InputField>
            <div className={styles.content}>
            <InputField type="date" title="Kdy?" placeholder="Zadejte datum..."></InputField>
            <SelectField title="V kolik?" options={["Ráno", "Odpoledne", "Večer"]} placeholder="Čas"></SelectField>
            </div>
        </>)
      }
      {
        phase === 1 && <div>
        <SelectField title="Typ" options={["Festivaly", "Divadla", "Kina", "Sport"]} placeholder="Typ"></SelectField>
       </div>
      }
      {
        phase === 2 && <div>
          Pozvat přátele
          <div className={styles.gap}></div>
          <div className={styles.content}>
        <img src="/assets/mockpfp.jpg" />
        <div className={styles.contentActions}>
          <h3>Matěj Tobiáš Moravec</h3>
          <a onClick={() => setInvited1(true)} className={styles.green}>{invited1 ? "pozvání odesláno" : "pozvat"}</a>
        </div> </div>
        <div className={styles.content}>
        <img src="/assets/mockpfp.jpg" />
        <div className={styles.contentActions}>
          <h3>Tomáš Kalhous</h3>
          <a onClick={() => setInvited2(true)} className={styles.green}>{invited2 ? "pozvání odesláno" : "pozvat"}</a>
        </div>
      </div>
        </div>
      }
      {
        phase === 3 && <div>
          <PlacesList></PlacesList>
        </div>
      }
     
        
      
        

       {phase === 2 ? <div className={styles.smGap}></div> : <div className={styles.gap}></div> } 
        <CustomButton
          size="small"
          onClick={phase !== 3 ? () => setPhase((phase) => phase + 1) : props.closeCallback}
          text={phase !== 3 ? "Pokračovat" : "Dokončit"}></CustomButton>
    </div>
  );
};

export default Meetup;
