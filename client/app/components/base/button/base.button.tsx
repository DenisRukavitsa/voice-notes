"use client";

interface Props {
  text: string;
  className?: string;
  icon?: JSX.Element;
  onClick?: () => void;
  disabled?: boolean;
}

const BaseButton = ({ text, className, icon, disabled, onClick }: Props) => {
  return (
    <button
      disabled={disabled}
      onClick={onClick}
      className={`${className} flex px-4 py-1 text-base font-semibold rounded-full border hover:text-white hover:border-transparent focus:outline-none disabled:bg-slate-50 disabled:text-slate-500 disabled:border-slate-200`}
    >
      {icon && icon}
      <span>{text}</span>
    </button>
  );
};

export default BaseButton;
