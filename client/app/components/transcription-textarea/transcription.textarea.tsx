import { useEffect, useState } from "react";

const TranscriptionTextarea = ({
  transcription,
}: {
  transcription: string;
}) => {
  const [textareaValue, setTextareaValue] = useState<string>("");

  useEffect(() => {
    setTextareaValue(transcription);
  }, [transcription]);

  return (
    <textarea
      rows={8}
      value={textareaValue}
      onChange={(e) => setTextareaValue(e.target.value)}
      placeholder="Transcribed note will go here..."
      className="mt-4 p-2 w-full rounded-md border border-slate-300 focus:border-gray-400 focus:outline-none"
    />
  );
};

export default TranscriptionTextarea;
