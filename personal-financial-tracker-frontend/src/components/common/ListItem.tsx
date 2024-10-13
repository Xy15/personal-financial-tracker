export interface ListItemProps {
  icon: string;
  title: string;
  description?: string;
  amount: number;
  time: string;
  styles?: {
    containerStyle?: React.CSSProperties;
  };
}

const ListItem = ({
  icon,
  title,
  description,
  amount,
  time,
  styles,
}: ListItemProps) => {
  return (
    <div
      className={`bg-background rounded-lg p-4 flex items-center m-2 ${styles?.containerStyle}`}
    >
      <img src={icon} className="max-w-[45px]" />
      <div className="px-3">
        <h1 className="font-bold">{title}</h1>
        <p className="truncate">
          {time} {description ?? ""}
        </p>
      </div>
      <div className="font-bold ml-auto mr-0">{amount.toFixed(2)}</div>
    </div>
  );
};

export default ListItem;
