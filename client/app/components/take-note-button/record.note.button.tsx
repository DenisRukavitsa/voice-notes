"use client";

import { MicrophoneIcon, StopIcon } from "@heroicons/react/24/outline";
import { useState } from "react";

interface Props {
  onRecordingStart: () => void;
  onRecordingStop: () => void;
}

const RecordNoteButton = ({ onRecordingStart, onRecordingStop }: Props) => {
  const [isRecording, setIsRecording] = useState(false);

  const handleRecordButtonClick = () => {
    setIsRecording(true);
    onRecordingStart();
  };

  const handleStopButtonClick = () => {
    setIsRecording(false);
    onRecordingStop();
  };

  if (isRecording) {
    return (
      <button
        onClick={handleStopButtonClick}
        className="flex mr-3 px-4 py-1 text-base font-semibold rounded-full border text-red-400 border-red-200 hover:text-white hover:bg-red-400 hover:border-transparent focus:outline-none"
      >
        <StopIcon className="size-5 mr-1" />
        <span>Stop recording</span>
      </button>
    );
  }

  return (
    <button
      onClick={handleRecordButtonClick}
      className="flex mr-3 px-4 py-1 text-base font-semibold rounded-full border text-teal-600 border-teal-300 hover:text-white hover:bg-teal-600 hover:border-transparent focus:outline-none"
    >
      <MicrophoneIcon className="size-5 mr-1" />
      <span>Record a note</span>
    </button>
  );
};

export default RecordNoteButton;
