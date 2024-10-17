import React from "react";
import InputField from "./InputField";
import styles from "./AuthDialog.module.css";
import CustomButton from "./CustomButton";
import { useNavigate } from "react-router-dom";

const AuthDialog = () => {
  const navigate = useNavigate();
  return (
    <div className={styles.center}>
      <div className={styles["justify-center"]}>
        <h1 className={styles.heading}>
          Friends<span className={styles.hk}>HK</span>
        </h1>
        <InputField type="email" placeholder="Email" title="Email" />
        <InputField type="password" placeholder="Heslo" title="Heslo" />
        <div className={styles.gap}></div>
        <CustomButton text={"Příhlásit se"} onClick={() => null}></CustomButton>
        <div className={styles.smallGap}></div>
        <CustomButton
          text={"Zaregistrovat"}
          onClick={() => navigate("/auth/register")}
          backgroundColor="#F0EFFF"
          textColor="#1B1937"
          hoverColor="#DEDEEC"
        ></CustomButton>
      </div>
    </div>
  );
};

export default AuthDialog;
