"use client";

import {
  ArrowPathIcon,
  MicrophoneIcon,
  StopIcon,
} from "@heroicons/react/24/outline";
import { useState } from "react";
import BaseButton from "../base/button/base.button";

interface Props {
  disabled?: boolean;
  loading?: boolean;
  onRecordingStart: () => Promise<boolean>;
  onRecordingStop: () => void;
}

const RecordNoteButton = ({
  disabled,
  loading,
  onRecordingStart,
  onRecordingStop,
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
      <BaseButton
        text={loading ? "Transcribing..." : "Stop recording"}
        icon={<StopIcon className="size-6 mr-1" />}
        disabled={disabled || loading}
        onClick={handleStopButtonClick}
        className="mr-3 text-red-400 border-red-200 hover:bg-red-400"
      />
    );
  }

  const icon = loading ? (
    <ArrowPathIcon className="size-6 mr-1 animate-spin" />
  ) : (
    <MicrophoneIcon className="size-5 mr-1" />
  );
  return (
    <BaseButton
      text={loading ? "Transcribing..." : "Record a note"}
      icon={icon}
      disabled={disabled || loading}
      onClick={handleRecordButtonClick}
      className="mr-3 text-teal-600 border-teal-300 hover:bg-teal-600"
    />
  );
};

export default RecordNoteButton;
