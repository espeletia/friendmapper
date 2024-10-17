import React from "react";

// Define styles for the input and icon
const styles = {
  inputContainer: {
    position: 'relative', // So that the icon can be positioned inside the input
    width: '100%', // Full width of the container
  } as React.CSSProperties,

  input: {
    width: '100%', // Full width input
    padding: '12px 20px', // Padding for input field
    paddingRight: '40px', // Add extra padding to the right to make space for the icon
    borderRadius: '30px', // Rounded corners
    border: 'none', // No border
    outline: 'none', // Remove the default outline
    fontSize: '16px', // Font size of input text
    transition: 'background-color 0.3s ease', // Smooth background transition
  } as React.CSSProperties,

  icon: {
    position: 'absolute', // Positioned inside the input
    right: '15px', // Distance from the right edge
    top: '50%', // Center vertically
    transform: 'translateY(-50%)', // Vertical alignment fix
    fontSize: '20px', // Icon size
    color: '#A0A0A0', // Default icon color
    cursor: 'pointer', // Pointer cursor on hover
  } as React.CSSProperties
};

interface SearchInputProps {

  placeholder: string;
  backgroundColor?: string;
  textColor?: string;
  iconColor?: string;
}

const SearchInput = (props: SearchInputProps) => {
  const {
    
    placeholder,
    backgroundColor = '#f0f0ff', // Default light background color
    textColor = '#888888', // Default grey placeholder text
    iconColor = '#A0A0A0' // Default grey icon color
  } = props;

  return (
    <div style={styles.inputContainer}>
      <input
        type="text"
        placeholder={placeholder}
        style={{
          ...styles.input,
          backgroundColor: backgroundColor, // Dynamic background
          color: textColor, // Dynamic text color
        }}
      />
      <span
        className="material-icons"
        style={{
          ...styles.icon,
          color: iconColor, // Dynamic icon color
        }}
      >
        search
      </span>
    </div>
  );
};

export default SearchInput;
