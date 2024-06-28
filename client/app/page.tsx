"use client";

import AudioTranscriber from "./components/audio-transcriber/audio.transcriber";
import CopyNoteButton from "./components/copy-note-button/copy.note.button";

export default function Home() {
  return (
    <main className="mx-auto p-4">
      <div className="flex">
        <AudioTranscriber onTranscription={() => {}}></AudioTranscriber>
        <CopyNoteButton></CopyNoteButton>
      </div>
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
