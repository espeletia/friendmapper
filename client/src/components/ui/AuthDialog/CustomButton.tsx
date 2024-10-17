import classNames from "classnames";
import React from "react";

interface ButtonProps {
  text: string;
  onClick: () => void;
  backgroundColor?: string;
  textColor?: string;
  hoverColor?: string; // Optional hover color
  variant?: "primary" | "secondary" | "third";
  size?: "small" | "large";
  className?: string;
}

const CustomButton = (props: ButtonProps) => {
  const [isHovered, setIsHovered] = React.useState(false);

  const {
    text,
    size = "large",
    variant = "primary",
    textColor = variant === "primary"
      ? "white"
      : variant === "secondary"
        ? "#EB1C1C"
        : variant === "third"
          ? "white"
          : "white", // Default to white text
    hoverColor = variant === "primary"
      ? "#5a56d1"
      : variant === "secondary"
        ? "#FFEDED"
        : variant === "third"
          ? "rgba(37, 194, 53, 1)"
          : "white", // Default hover color (darker purple)
    backgroundColor = variant === "primary"
      ? "#6C63FF"
      : variant === "secondary"
        ? "#FFEDED"
        : variant === "third"
          ? "rgba(37, 194, 53, 1)"
          : "white", // Default to purple if no backgroundColor prop is passed
    onClick,
  } = props;

  const styles = {
    button: {
      width: "100%", // Full width button
      padding: size === "large" ? "12px 30px" : "12px 12px", // Padding inside the button
      border: "none", // Remove default border
      borderRadius: "30px", // Rounded corners
      fontSize: "16px", // Text size
      fontWeight: "bold", // Bold text
      textAlign: "center", // Center the text
      cursor: "pointer", // Pointer cursor on hover
      transition: "background-color 0.3s ease", // Smooth transition on hover
      height: size === "large" ? "auto" : "2.75rem", // Set height based on size prop
      ...classNames,
    } as React.CSSProperties,

    buttonHover: {
      // Default hover color will be applied dynamically if needed
    } as React.CSSProperties,
  };

  return (
    <button
      style={{
        ...styles.button,
        backgroundColor: isHovered ? hoverColor : backgroundColor,
        color: textColor,
      }}
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
      onClick={onClick}
    >
      {text}
    </button>
  );
};

export default CustomButton;
