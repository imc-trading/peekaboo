class Peekaboo < Formula
  version "0.2.1"
  homepage "https://github.com/mickep76/peekaboo"
  url "https://github.com/mickep76/peekaboo/archive/#{version}.tar.gz"
  sha256 "3ad8d25fafbbb730509b49bef94be4994d58c4725efd0189eab3cf959e6a3324"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    system "./build"
    bin.install "bin/peekaboo"
    (prefix/"peekaboo").install "src/github.com/mickep76/peekaboo/static", "src/github.com/mickep76/peekaboo/templates"
  end

  def plist; <<-EOS.undent
    <?xml version="1.0" encoding="UTF-8"?>
    <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
    <plist version="1.0">
    <dict>
      <key>KeepAlive</key>
      <true/>
      <key>Label</key>
      <string>#{plist_name}</string>
      <key>ProgramArguments</key>
      <array>
        <string>#{bin}/peekaboo</string>
        <string>--static-dir</string>
        <string>#{prefix}/peekaboo/static</string>
        <string>--template-dir</string>
        <string>#{prefix}/peekaboo/templates</string>
      </array>
      <key>RunAtLoad</key>
      <true/>
    </dict>
    </plist>
    EOS
  end

  test do
    system "#{bin}/peekaboo", "--version"
  end
end
