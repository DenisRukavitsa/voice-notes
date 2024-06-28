class AudioRecorder {
  private mediaRecorder: MediaRecorder;
  private audioChunks: Blob[] = [];
  private listeners: Array<(audioBlob: Blob) => void> = [];

  async start() {
    if (!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia) {
      console.error('getUserMedia is not supported');
      return;
    }

    const mediaStream = await navigator.mediaDevices.getUserMedia({ audio: true });
    this.mediaRecorder = new MediaRecorder(mediaStream);
    this.mediaRecorder.ondataavailable = (event) => this.audioChunks.push(event.data);
    this.mediaRecorder.onstop = () => this.notifyListeners();
    this.mediaRecorder.start();
  }

  stop() {
    this.mediaRecorder?.stop();
    this.mediaRecorder?.stream.getTracks().forEach((track) => track.stop());
  }

  onData(callback: (audioBlob: Blob) => void) {
    this.listeners.push(callback);
  }

  private createAudioBlob(): Blob {
    const audioBlob = new Blob(this.audioChunks, { type: 'audio/ogg; codecs=opus' });
    this.audioChunks = [];
    return audioBlob;
  }

  private notifyListeners() {
    const audioBlob = this.createAudioBlob();
    this.listeners.forEach((listener) => listener(audioBlob));
  }
}

export default AudioRecorder;
