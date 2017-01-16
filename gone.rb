class Gone < Formula
  desc "Micro pomodoro CLI"
  homepage "https://github.com/guillaumebreton/gone"
  url "https://github.com/guillaumebreton/gone/releases/download/2.3.1/gone_Darwin_x86_64.tar.gz"
  version "2.3.1"
  sha256 "1b7bbfe13b3c85971b1c0882f623892d9c1f35800f1d1e7c25e114b8285752e9"

  def install
    bin.install "gone"
  end
end
