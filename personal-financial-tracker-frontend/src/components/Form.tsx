import { useForm, FormProvider } from "react-hook-form";
import Input, { InputProps } from "./Input";
import Button from "./Button";

interface FormProps {
  formInputs: InputProps[];
  onSubmit: (data: any) => void;
}

const Form = ({formInputs, onSubmit}:FormProps) => {
  const methods = useForm();
  const { handleSubmit } = methods;

  return (
    <FormProvider {...methods}>
      <div className="bg-border rounded-lg w-1/3 p-1.5">
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="bg-background rounded border-background p-5 w-full"
        >
          {formInputs.map((input) => {
            return <Input key={input.name} {...input} />;
          })}
          <div className="flex justify-center mt-4">
            <Button type="submit">Submit</Button>
          </div>
        </form>
      </div>
    </FormProvider>
  );
};

export default Form;
