import React from "react";

const styles = {
  button: {
    width: '100%', // Full width button
    padding: '12px 20px', // Padding inside the button
    border: 'none', // Remove default border
    borderRadius: '30px', // Rounded corners
    fontSize: '16px', // Text size
    fontWeight: 'bold', // Bold text
    textAlign: 'center', // Center the text
    cursor: 'pointer', // Pointer cursor on hover
    transition: 'background-color 0.3s ease', // Smooth transition on hover
  } as React.CSSProperties,

  buttonHover: {
    // Default hover color will be applied dynamically if needed
  } as React.CSSProperties
};

interface ButtonProps {
  text: string;
  onClick: () => void;
  backgroundColor?: string;
  textColor?: string;
  hoverColor?: string; // Optional hover color
}

const CustomButton = (props: ButtonProps) => {
  const [isHovered, setIsHovered] = React.useState(false);

  const {
    text, 
    backgroundColor = '#6C63FF', // Default to purple if no backgroundColor prop is passed
    textColor = 'white', // Default to white text
    hoverColor = '#5a56d1', // Default hover color (darker purple)
    onClick
  } = props;

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
