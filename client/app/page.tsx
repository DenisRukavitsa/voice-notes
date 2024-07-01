"use client";

import { useState } from "react";
import AudioTranscriber from "./components/audio-transcriber/audio.transcriber";
import CopyNoteButton from "./components/copy-note-button/copy.note.button";
import ErrorMessage from "./components/error-message/error.message";
import TranscriptionTextarea from "./components/transcription-textarea/transcription.textarea";

export default function Home() {
  const [error, setError] = useState<string>("");
  const [transcription, setTranscription] = useState<string>("");

  return (
    <main className="mx-auto p-4">
      <div className="flex">
        <AudioTranscriber
          onTranscription={(transcription) => setTranscription(transcription)}
          onError={(error) => setError(error)}
          clearError={() => setError("")}
        />
        <CopyNoteButton />
      </div>
      {error && <ErrorMessage error={error} />}
      <TranscriptionTextarea transcription={transcription} />
    </main>
  );
}
