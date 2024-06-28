"use client";

import AudioRecorder from "@/app/utils/audio-recorder/audio.recorder";
import { useEffect, useState } from "react";
import RecordNoteButton from "../take-note-button/record.note.button";

interface Props {
  onTranscription: (transcription: string) => void;
}

const AudioTranscriber = ({ onTranscription }: Props) => {
  const [isAudioInputAvailable, setIsAudioInputAvailable] = useState(true);
  const [isTranscribing, setIsTranscribing] = useState(false);
  const [audioRecorder, setAudioRecorder] = useState<AudioRecorder | null>(
    null
  );

  const handleAudioRecorderData = async (blob: Blob) => {
    setIsTranscribing(true);
    try {
      const audioFile = new File([blob], "audio.ogg", { type: "audio/ogg" });
      console.log(audioFile);
      // TODO: send file to the backend
    } catch (error) {
      console.error(error);
    } finally {
      setIsTranscribing(false);
    }
  };

  useEffect(() => {
    if (!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia) {
      console.error("getUserMedia is not supported");
      setIsAudioInputAvailable(false);
      return;
    }

    const recorder = new AudioRecorder();
    setAudioRecorder(recorder);
    recorder.onData(handleAudioRecorderData);
  }, []);

  const handleRecordingStart = async () => {
    if (!audioRecorder) return;
    try {
      await audioRecorder.start();
    } catch (error) {
      console.error(error);
    }
  };

  const handleRecordingStop = async () => {
    if (!audioRecorder) return;
    audioRecorder.stop();
  };

  return (
    <RecordNoteButton
      onRecordingStart={handleRecordingStart}
      onRecordingStop={handleRecordingStop}
    ></RecordNoteButton>
  );
};

export default AudioTranscriber;
