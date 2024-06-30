"use client";

import { useState } from "react";
import AudioTranscriber from "./components/audio-transcriber/audio.transcriber";
import CopyNoteButton from "./components/copy-note-button/copy.note.button";
import ErrorMessage from "./components/error-message/error.message";

export default function Home() {
  const [error, setError] = useState<string>("");

  const handleError = (error: string) => {
    setError(error);
  };

  return (
    <main className="mx-auto p-4">
      <div className="flex">
        <AudioTranscriber
          onTranscription={() => {}}
          onError={handleError}
        ></AudioTranscriber>
        <CopyNoteButton></CopyNoteButton>
      </div>

      {error && <ErrorMessage error={error}></ErrorMessage>}

      <div>
        <textarea
          rows={8}
          placeholder="Transcribed note will go here..."
          className="mt-4 p-2 w-full rounded-md border border-slate-300 focus:border-gray-400 focus:outline-none"
        ></textarea>
      </div>
    </main>
  );
}
