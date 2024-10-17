import React from "react";
import InputField from "./InputField";
import styles from "./AuthDialog.module.css";
import CustomButton from "./CustomButton";


const RegisterDialog = () => {
  return (
    <div>
      <div className={styles['justify-center']}>
        <h1 className={styles.heading}>
          Friends<span className={styles.hk}>HK</span>
        </h1>
        <InputField type="text" placeholder="Zadejte vaše jméno..." title="Jméno"/>
        <InputField type="text" placeholder="@blackmamba" title="Jméno v aplikaci"/>
        <InputField type="email" placeholder="Zadejte email" title="Email"/>
        <InputField type="password" placeholder="Zadejte heslo..." title= "Heslo"/>
        <InputField type="password" placeholder="Zadejte heslo..." title= "Heslo (znovu)"/>
        <div className={styles.gap}></div>
        <CustomButton text={"Zaregistrovat se"} onClick={() => null}></CustomButton>
        <div className={styles.smallGap}></div>
        
      </div>
    </div>
  );
};

export default RegisterDialog;

