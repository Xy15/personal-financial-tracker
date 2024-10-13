import { Controller, useFormContext, RegisterOptions } from "react-hook-form";

export interface SelectOptionsProps {
  name: string;
  value: string;
}

export interface InputProps {
  name: string;
  label: string;
  type?: string;
  options?: SelectOptionsProps[];
  className?: {
    container?: string;
    label?: string;
    input?: string;
    option?: string;
  };
  required?: boolean;
  rules?: RegisterOptions;
}

const Input = ({
  name,
  label,
  type = "text",
  options,
  className,
  required = true,
  rules,
}: InputProps) => {
  const { control } = useFormContext();

  return (
    <Controller
      name={name}
      defaultValue={type == "select" && options ? options[0].value : ""}
      control={control}
      rules={{
        required: required
          ? {
              value: true,
              message: label + " is required",
            }
          : false,
        ...rules,
      }}
      render={({ field, fieldState }) => (
        <div
          className={`grid grid-cols-3 p-2 items-center ${className?.container}`}
        >
          <label htmlFor={name} className={`col-span-1 ${className?.label}`}>
            {label}
          </label>
          {type == "select" ? (
            <select
              {...field}
              id={name}
              className={`col-span-2 p-2 ${
                className?.input && className?.input
              }`}
            >
              {options?.map((option) => (
                <option
                  key={option.value}
                  value={option.value}
                  className={`${className?.input}`}
                >
                  {option.name}
                </option>
              ))}
            </select>
          ) : (
            <input
              {...field}
              id={name}
              type={type}
              className={`col-span-2 p-2 ${className?.input}`}
            />
          )}
          {fieldState.error && (
            <>
              <div></div>
              <span className="col-span-2 text-sm text-danger">
                {fieldState.error.message}
              </span>
            </>
          )}
        </div>
      )}
    />
  );
};

export default Input;
