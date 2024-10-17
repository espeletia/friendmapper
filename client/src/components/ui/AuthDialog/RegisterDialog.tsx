import React from "react";
import InputField from "./InputField";
import styles from "./AuthDialog.module.css";
import CustomButton from "./CustomButton";
import ProfilePicturePlaceholder from "./ProfilePicturePlaceholder";
import { useNavigate } from "react-router-dom";

const RegisterDialog = () => {
  const navigate = useNavigate();
  return (
    <div className={styles.center}>
      <div className={styles["justify-center"]}>
        <h1 className={styles.heading}>
          Friends<span className={styles.hk}>HK</span>
        </h1>
        <ProfilePicturePlaceholder />
        <InputField
          type="text"
          placeholder="Zadejte vaše jméno..."
          title="Jméno"
        />
        <InputField type="email" placeholder="Zadejte email" title="Email" />
        <InputField
          type="password"
          placeholder="Zadejte heslo..."
          title="Heslo"
        />
        <InputField
          type="password"
          placeholder="Zadejte heslo..."
          title="Heslo (znovu)"
        />
        <div className={styles.gap}></div>
        <CustomButton
          text={"Zaregistrovat se"}
          onClick={() => navigate("/map")}
        ></CustomButton>
        <div className={styles.smallGap}></div>
      </div>
    </div>
  );
};

export default RegisterDialog;
