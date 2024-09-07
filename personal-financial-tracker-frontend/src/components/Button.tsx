import { ReactNode } from "react";

interface ButtonProps {
  type?: "submit" | "reset" | "button";
  onClick?: React.MouseEventHandler<HTMLButtonElement>;
  disabled?: boolean;
  className?: string;
  children: ReactNode;
}

const Button = ({
  type,
  onClick,
  disabled = false,
  className,
  children,
}: ButtonProps) => {
  return (
    <button type={type} onClick={onClick} disabled={disabled} className={`bg-secondary hover:opacity-80 active:opacity-60 px-6 py-2 rounded-lg ${className}`}>
      {children}
    </button>
  )
}

export default Button