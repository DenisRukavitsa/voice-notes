"use client";

import { MicrophoneIcon, StopIcon } from "@heroicons/react/24/outline";
import { useState } from "react";

interface Props {
  disabled?: boolean;
  onRecordingStart: () => Promise<boolean>;
  onRecordingStop: () => void;
}

const RecordNoteButton = ({
  onRecordingStart,
  onRecordingStop,
  disabled,
}: Props) => {
  const [isRecording, setIsRecording] = useState(false);

  const handleRecordButtonClick = async () => {
    const recordingStarted = await onRecordingStart();
    if (!recordingStarted) return;
    setIsRecording(true);
  };

  const handleStopButtonClick = () => {
    setIsRecording(false);
    onRecordingStop();
  };

  if (isRecording) {
    return (
      <button
        disabled={disabled}
        onClick={handleStopButtonClick}
        className="flex mr-3 px-4 py-1 text-base font-semibold rounded-full border text-red-400 border-red-200 hover:text-white hover:bg-red-400 hover:border-transparent focus:outline-none disabled:bg-slate-50 disabled:text-slate-500 disabled:border-slate-200"
      >
        <StopIcon className="size-5 mr-1" />
        <span>Stop recording</span>
      </button>
    );
  }

  return (
    <button
      disabled={disabled}
      onClick={handleRecordButtonClick}
      className="flex mr-3 px-4 py-1 text-base font-semibold rounded-full border text-teal-600 border-teal-300 hover:text-white hover:bg-teal-600 hover:border-transparent focus:outline-none disabled:bg-slate-50 disabled:text-slate-500 disabled:border-slate-200"
    >
      <MicrophoneIcon className="size-5 mr-1" />
      <span>Record a note</span>
    </button>
  );
};

export default RecordNoteButton;
