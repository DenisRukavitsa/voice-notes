"use client";

import {
  CheckCircleIcon,
  DocumentDuplicateIcon,
} from "@heroicons/react/24/outline";
import { useEffect, useState } from "react";

export default function CopyNoteButton({ noteText }: { noteText: string }) {
  const [isCopied, setIsCopied] = useState(false);

  useEffect(() => {
    if (isCopied) {
      const timeout = setTimeout(() => {
        setIsCopied(false);
      }, 2000);

      return () => clearTimeout(timeout);
    }
  }, [isCopied]);

  const handleNoteCopy = () => {
    navigator.clipboard.writeText(noteText);
    setIsCopied(true);
  };

  return (
    <button
      onClick={handleNoteCopy}
      disabled={!noteText}
      className="flex px-4 py-1 text-base text-teal-600 font-semibold rounded-full border border-teal-300 hover:text-white hover:bg-teal-600 hover:border-transparent focus:outline-none disabled:bg-slate-50 disabled:text-slate-500 disabled:border-slate-200"
    >
      {isCopied ? (
        <CheckCircleIcon className="size-6 mr-1" />
      ) : (
        <DocumentDuplicateIcon className="size-5 mr-1" />
      )}
      {isCopied ? <span>Copied!</span> : <span>Copy note</span>}
    </button>
  );
}
