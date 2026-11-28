import './home.css';

import { AgentSkillCard, InstallBanner, LandingHero, LandingSection } from '@lobehub/ui/awesome';
import { GithubIcon } from '@lobehub/ui/icons';
import { ArrowRight } from 'lucide-react';
import { Link, useNavigate } from 'react-router';

import { Features } from './Features';
import { Gallery } from './Gallery';
import { renderLink } from './renderLink';
import { ShowcaseFrame } from './ShowcaseFrame';
import { Snippets } from './Snippets';

const GITHUB_URL = 'https://github.com/puremixlab/purekit';
const GET_STARTED_SOURCE = GITHUB_URL + '/blob/main/docs/getting-started.mdx';
const INSTALL_COMMAND = 'go get github.com/puremixlab/purekit';
const GET_STARTED = '/getting-started';

export default function Home({ description }: { description: string }) {
  const navigate = useNavigate();
  return (
    <>
      <LandingHero
        accent={<span className="purekit-accent">Native Go UI</span>}
        description={<span data-pagefind-meta="description">{description}</span>}
        renderLink={renderLink}
        title={<span data-pagefind-meta="title">PureKit</span>}
        actions={[
          { href: GET_STARTED, icon: ArrowRight, iconPlacement: 'end', label: 'Get Started', primary: true },
          { href: GITHUB_URL, icon: GithubIcon, label: 'GitHub' },
        ]}
        aside={
          <AgentSkillCard
            agent={{
              code: `Read ${GET_STARTED_SOURCE} and use PureKit to build the native Go UI.`,
              description: 'Send this prompt to your agent to pick the right components',
            }}
            footer={
              <>
                <a href="/skills.md" rel="noreferrer" target="_blank">
                  skills.md
                </a>
                <a href="/llms.txt" rel="noreferrer" target="_blank">
                  llms.txt
                </a>
              </>
            }
            human={{
              code: INSTALL_COMMAND,
              description: 'Add PureKit to your Go module',
            }}
          />
        }
        badge={
          <a href="https://github.com/egoist/mygo" rel="noreferrer" target="_blank">
            Built on MyGo · GPU-drawn native UI →
          </a>
        }
        onNavigate={navigate}
      >
        <ShowcaseFrame />
      </LandingHero>

      <LandingSection
        actions={[{ href: '/components/agent-avatar', icon: ArrowRight, iconPlacement: 'end', label: 'Browse components' }]}
        description="Each tile previews a Go component in light or dark mode."
        eyebrow="Components"
        eyebrowColor="blue"
        id="home-gallery"
        renderLink={renderLink}
        title="Explore the components"
        onNavigate={navigate}
      >
        <Gallery />
      </LandingSection>

      <LandingSection
        actions={[{ href: GET_STARTED, label: 'Getting started' }]}
        description="Immediate-mode Go: a view is a function, state is your own variables."
        eyebrow="Developer experience"
        eyebrowColor="green"
        id="home-snippets"
        renderLink={renderLink}
        title="Code you write, UI you get"
        onNavigate={navigate}
      >
        <Snippets />
      </LandingSection>

      <LandingSection
        actions={[{ href: '/components/tokens', label: 'Design tokens' }]}
        description="The foundations a desktop app needs, shared with the web."
        eyebrow="Foundations"
        eyebrowColor="orange"
        id="home-features"
        renderLink={renderLink}
        title="One design language, two runtimes"
        onNavigate={navigate}
      >
        <Features />
      </LandingSection>

      <InstallBanner
        command={INSTALL_COMMAND}
        prefix="$"
        title="Build your native Go app"
        footnote={
          <>
            Open source · MIT license · Go 1.27.3+ · <Link to={GET_STARTED}>Get Started →</Link>
          </>
        }
      />
    </>
  );
}
