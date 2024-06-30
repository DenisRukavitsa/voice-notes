"use client";

import AudioRecorder from "@/app/utils/audio-recorder/audio.recorder";
import { useEffect, useState } from "react";
import RecordNoteButton from "../record-note-button/record.note.button";

interface Props {
  onTranscription: (transcription: string) => void;
  onError: (error: string) => void;
}

const AudioTranscriber = ({ onTranscription, onError }: Props) => {
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
      onError(`Error transcribing audio: ${error}`);
    } finally {
      setIsTranscribing(false);
    }
  };

  useEffect(() => {
    if (!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia) {
      onError("Audio input is not available on this device");
      setIsAudioInputAvailable(false);
      return;
    }

    const recorder = new AudioRecorder();
    setAudioRecorder(recorder);
    recorder.onData(handleAudioRecorderData);
  }, []);

  const handleRecordingStart = async () => {
    if (!audioRecorder) return false;
    try {
      await audioRecorder.start();
    } catch (error) {
      console.error(error);
      onError(`Error recording audio: ${error}`);
      return false;
    }

    return true;
  };

  const handleRecordingStop = () => {
    if (!audioRecorder) return;
    audioRecorder.stop();
  };

  return (
    <RecordNoteButton
      disabled={!isAudioInputAvailable}
      onRecordingStart={handleRecordingStart}
      onRecordingStop={handleRecordingStop}
    />
  );
};

export default AudioTranscriber;
