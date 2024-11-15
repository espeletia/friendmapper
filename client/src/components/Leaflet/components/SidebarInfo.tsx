import { Cross, CrossBlack } from "../../../svg";
import { Marker } from "../../../types";
import CustomButton from "../../ui/AuthDialog/CustomButton";

import styles from "./SidebarInfo.module.css";

interface Props {
  marker: Marker;
  closeCallback: VoidFunction;
}

const SidebarInfo = (props: Props) => {
  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <div className={styles.header}>
          <h1>{props.marker.Name}</h1>
          <button onClick={props.closeCallback}>
            <CrossBlack />
          </button>
        </div>
        <p>{props.marker.Address}</p>
        {props.marker.Accessibility ? (
          <CustomButton
            text="Bezbariérový přístup"
            size="small"
            variant="third"
            onClick={() => {}}
          />
        ) : (
          <></>
        )}
        <p>{props.marker.Description}</p>
        <div className={styles.buttons}>
          <CustomButton
            onClick={() => {}}
            variant="primary"
            size="small"
            text="Program"
          ></CustomButton>
          <CustomButton
            onClick={() => {}}
            variant="primary"
            size="small"
            text="Vytvořit meetup"
          ></CustomButton>
          <CustomButton
            onClick={() => {}}
            variant="primary"
            size="small"
            text="Like"
          ></CustomButton>
          <CustomButton
            onClick={() => {}}
            variant="primary"
            size="small"
            text="Webové stránky"
          ></CustomButton>
        </div>
      </div>
    </div>
  );
};

export default SidebarInfo;
