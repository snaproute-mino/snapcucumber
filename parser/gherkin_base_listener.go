// Generated from Gherkin.g4 by ANTLR 4.7.

package parser // Gherkin

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
)

// BaseGherkinListener is a complete listener for a parse tree produced by GherkinParser.
type BaseGherkinListener struct{}

var _ GherkinListener = &BaseGherkinListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGherkinListener) VisitTerminal(node antlr.TerminalNode) { log.Print("VisitTerminal") }

// VisitErrorNode is called when an error node is visited.
func (s *BaseGherkinListener) VisitErrorNode(node antlr.ErrorNode) { log.Print("VisitError") }

// EnterEveryRule is called when any rule is entered.
func (s *BaseGherkinListener) EnterEveryRule(ctx antlr.ParserRuleContext) { log.Print("EnterEveryRule") }

// ExitEveryRule is called when any rule is exited.
func (s *BaseGherkinListener) ExitEveryRule(ctx antlr.ParserRuleContext) { log.Print("ExitEveryRule") }

// EnterTagName is called when production tagName is entered.
func (s *BaseGherkinListener) EnterTagName(ctx *TagNameContext) { log.Print("EnterTagName") }

// ExitTagName is called when production tagName is exited.
func (s *BaseGherkinListener) ExitTagName(ctx *TagNameContext) { log.Print("ExitTagName") }

// EnterTag is called when production tag is entered.
func (s *BaseGherkinListener) EnterTag(ctx *TagContext) { log.Print("EnterTag") }

// ExitTag is called when production tag is exited.
func (s *BaseGherkinListener) ExitTag(ctx *TagContext) { log.Print("ExitTag") }

// EnterLineContent is called when production lineContent is entered.
func (s *BaseGherkinListener) EnterLineContent(ctx *LineContentContext) { log.Print("EnterLineContent") }

// ExitLineContent is called when production lineContent is exited.
func (s *BaseGherkinListener) ExitLineContent(ctx *LineContentContext) { log.Print("ExitLineContent") }

// EnterComment is called when production comment is entered.
func (s *BaseGherkinListener) EnterComment(ctx *CommentContext) { log.Print("EnterComment") }

// ExitComment is called when production comment is exited.
func (s *BaseGherkinListener) ExitComment(ctx *CommentContext) { log.Print("ExitComment") }

// EnterAnnotation is called when production annotation is entered.
func (s *BaseGherkinListener) EnterAnnotation(ctx *AnnotationContext) { log.Print("EnterAnnotation") }

// ExitAnnotation is called when production annotation is exited.
func (s *BaseGherkinListener) ExitAnnotation(ctx *AnnotationContext) { log.Print("ExitAnnotation") }

// EnterTrailingComment is called when production trailingComment is entered.
func (s *BaseGherkinListener) EnterTrailingComment(ctx *TrailingCommentContext) {
	log.Print("EnterTrailingComment")
}

// ExitTrailingComment is called when production trailingComment is exited.
func (s *BaseGherkinListener) ExitTrailingComment(ctx *TrailingCommentContext) {
	log.Print("ExitTrailingComment")
}

// EnterDescriptionLine is called when production descriptionLine is entered.
func (s *BaseGherkinListener) EnterDescriptionLine(ctx *DescriptionLineContext) {
	log.Print("EnterDescriptionLine")
}

// ExitDescriptionLine is called when production descriptionLine is exited.
func (s *BaseGherkinListener) ExitDescriptionLine(ctx *DescriptionLineContext) {
	log.Print("ExitDescriptionLine")
}

// EnterName is called when production name is entered.
func (s *BaseGherkinListener) EnterName(ctx *NameContext) { log.Print("EnterName") }

// ExitName is called when production name is exited.
func (s *BaseGherkinListener) ExitName(ctx *NameContext) { log.Print("ExitName") }

// EnterFeature is called when production feature is entered.
func (s *BaseGherkinListener) EnterFeature(ctx *FeatureContext) { log.Print("EnterFeature") }

// ExitFeature is called when production feature is exited.
func (s *BaseGherkinListener) ExitFeature(ctx *FeatureContext) { log.Print("ExitFeature") }

// EnterBackground is called when production background is entered.
func (s *BaseGherkinListener) EnterBackground(ctx *BackgroundContext) { log.Print("EnterBackground") }

// ExitBackground is called when production background is exited.
func (s *BaseGherkinListener) ExitBackground(ctx *BackgroundContext) { log.Print("ExitBackground") }

// EnterAbstractScenario is called when production abstractScenario is entered.
func (s *BaseGherkinListener) EnterAbstractScenario(ctx *AbstractScenarioContext) {
	log.Print("EnterAbstractScenario")
}

// ExitAbstractScenario is called when production abstractScenario is exited.
func (s *BaseGherkinListener) ExitAbstractScenario(ctx *AbstractScenarioContext) {
	log.Print("ExitAbstractScenario")
}

// EnterScenario is called when production scenario is entered.
func (s *BaseGherkinListener) EnterScenario(ctx *ScenarioContext) { log.Print("EnterScenario") }

// ExitScenario is called when production scenario is exited.
func (s *BaseGherkinListener) ExitScenario(ctx *ScenarioContext) { log.Print("ExitScenario") }

// EnterOutline is called when production outline is entered.
func (s *BaseGherkinListener) EnterOutline(ctx *OutlineContext) { log.Print("EnterOutline") }

// ExitOutline is called when production outline is exited.
func (s *BaseGherkinListener) ExitOutline(ctx *OutlineContext) { log.Print("ExitOutline") }

// EnterExamples is called when production examples is entered.
func (s *BaseGherkinListener) EnterExamples(ctx *ExamplesContext) { log.Print("EnterExamples") }

// ExitExamples is called when production examples is exited.
func (s *BaseGherkinListener) ExitExamples(ctx *ExamplesContext) { log.Print("ExitExamples") }

// EnterStep is called when production step is entered.
func (s *BaseGherkinListener) EnterStep(ctx *StepContext) { log.Print("EnterStep") }

// ExitStep is called when production step is exited.
func (s *BaseGherkinListener) ExitStep(ctx *StepContext) { log.Print("ExitStep") }

// EnterRow is called when production row is entered.
func (s *BaseGherkinListener) EnterRow(ctx *RowContext) { log.Print("EnterRow") }

// ExitRow is called when production row is exited.
func (s *BaseGherkinListener) ExitRow(ctx *RowContext) { log.Print("ExitRow") }

// EnterCell is called when production cell is entered.
func (s *BaseGherkinListener) EnterCell(ctx *CellContext) { log.Print("EnterCell") }

// ExitCell is called when production cell is exited.
func (s *BaseGherkinListener) ExitCell(ctx *CellContext) { log.Print("ExitCell") }

// EnterDocument is called when production document is entered.
func (s *BaseGherkinListener) EnterDocument(ctx *DocumentContext) { log.Print("EnterDocument") }

// ExitDocument is called when production document is exited.
func (s *BaseGherkinListener) ExitDocument(ctx *DocumentContext) { log.Print("ExitDocument") }

// EnterDocumentIndent is called when production documentIndent is entered.
func (s *BaseGherkinListener) EnterDocumentIndent(ctx *DocumentIndentContext) {
	log.Print("EnterDocumentIndent")
}

// ExitDocumentIndent is called when production documentIndent is exited.
func (s *BaseGherkinListener) ExitDocumentIndent(ctx *DocumentIndentContext) {
	log.Print("ExitDocumentIndent")
}

// EnterDocumentContent is called when production documentContent is entered.
func (s *BaseGherkinListener) EnterDocumentContent(ctx *DocumentContentContext) {
	log.Print("EnterDocumentContent")
}

// ExitDocumentContent is called when production documentContent is exited.
func (s *BaseGherkinListener) ExitDocumentContent(ctx *DocumentContentContext) {
	log.Print("ExitDocumentContent")
}
