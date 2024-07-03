import { ChangeEvent, useEffect, useState } from "react";

interface Props {
  transcription: string;
  onChange: (transcription: string) => void;
}

const TranscriptionTextarea = ({ transcription, onChange }: Props) => {
  const [textareaValue, setTextareaValue] = useState<string>("");

  useEffect(() => {
    setTextareaValue(transcription);
  }, [transcription]);

  const handleChange = (event: ChangeEvent<HTMLTextAreaElement>) => {
    setTextareaValue(event.target.value);
    onChange(event.target.value);
  };

  return (
    <textarea
      rows={8}
      value={textareaValue}
      onChange={(event) => handleChange(event)}
      placeholder="Transcribed note will go here..."
      className="mt-4 p-2 w-full rounded-md border border-slate-300 focus:border-gray-400 focus:outline-none"
    />
  );
};

export default TranscriptionTextarea;
