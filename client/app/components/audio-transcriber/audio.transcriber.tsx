"use client";

import AudioRecorder from "@/app/utils/audio-recorder/audio.recorder";
import { useCallback, useEffect, useState } from "react";
import RecordNoteButton from "../record-note-button/record.note.button";

interface Props {
  onTranscription: (transcription: string) => void;
  onError: (error: string) => void;
  clearError: () => void;
}

const AudioTranscriber = ({ onTranscription, onError, clearError }: Props) => {
  const [isAudioInputAvailable, setIsAudioInputAvailable] = useState(true);
  const [isTranscribing, setIsTranscribing] = useState(false);
  const [audioRecorder, setAudioRecorder] = useState<AudioRecorder | null>(
    null
  );

  const handleAudioRecorderData = useCallback(
    async (blob: Blob) => {
      setIsTranscribing(true);
      try {
        const transcription = await transcribeAudio(blob);
        onTranscription(transcription);
      } catch (error) {
        console.error(error);
        onError(`Error transcribing audio: ${error}`);
      } finally {
        setIsTranscribing(false);
      }
    },
    [onTranscription, onError]
  );

  useEffect(() => {
    if (!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia) {
      onError("Audio input is not available on this device");
      setIsAudioInputAvailable(false);
      return;
    }

    const recorder = new AudioRecorder();
    setAudioRecorder(recorder);
    recorder.onData(handleAudioRecorderData);
  }, [onError, handleAudioRecorderData]);

  const transcribeAudio = async (audioBlob: Blob) => {
    const audioFile = new File([audioBlob], "audio.mp4", { type: "audio/mp4" });
    const formData = new FormData();
    formData.append("file", audioFile);

    const apiUrl = `${process.env.NEXT_PUBLIC_API_URL}/transcribe`;
    const response = await fetch(apiUrl, {
      method: "POST",
      body: formData,
    });
    const json = await response.json();

    if (!response.ok) {
      throw new Error(json.error);
    }

    return json.transcription;
  };

  const handleRecordingStart = async () => {
    if (!audioRecorder) return false;

    clearError();
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
      loading={isTranscribing}
      onRecordingStart={handleRecordingStart}
      onRecordingStop={handleRecordingStop}
    />
  );
};

export default AudioTranscriber;
