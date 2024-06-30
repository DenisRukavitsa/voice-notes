const ErrorMessage = ({ error }: { error: string }) => {
  return (
    <div className="mt-2 text-sm text-red-400 italic">
      <p>{error}</p>
    </div>
  );
};

export default ErrorMessage;
